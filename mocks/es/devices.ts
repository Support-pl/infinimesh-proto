import { v4 as uuidv4 } from 'uuid'
import { JsonValue, Struct } from '@bufbuild/protobuf'
import { createPromiseClient, createRouterTransport, ConnectError, Code } from '@connectrpc/connect'
import { DevicesService, NamespacesService } from '../../build/es/node/node_connect'
import { DeleteResponse, EmptyMessage, TokenResponse } from '../../build/es//node/node_pb'
import { Access, Nodes, Node, Role } from '../../build/es/node/access/access_pb'
import { Certificate, CreateResponse, Device, Devices } from '../../build/es/node/devices/devices_pb'
import { transport as nsTransport } from './namespaces.ts'

export const transport = createRouterTransport(({ service }) => {
  const namespaceApi = createPromiseClient(NamespacesService, nsTransport)
  const devices = new Map<string, Device>()
  const devicesByNs = new Map<string, Device[]>()
  const nodes = new Map<string, Node[]>()

  function changeDevice(key: keyof Device, uuid: string, value?: JsonValue) {
    const device = devices.get(uuid) ?? new Device({ uuid })

    if (typeof value !== 'boolean' && (value ?? true)) {
      value = !device[key]
    }

    for (const devices of devicesByNs.values()) {
      const device = devices.find(({ uuid: id }) => id === uuid)

      if (device) {
        (device[key] as JsonValue | undefined) = value
        break
      }
    }

    (device[key] as JsonValue | undefined) = value
    return device
  }

  namespaceApi.list(new EmptyMessage()).then(({ namespaces }) => {
    const namespacesList = namespaces.map(({ uuid }) => uuid)

    for (const key of namespacesList) {
      const length = Math.floor(Math.random() * 10)

      devicesByNs.set(key, Array.from({ length }).map((_, i) => {
        const value = Math.random()
        const length = Math.floor(value * 6)
        const role = Math.floor(Math.random() * 2 + 1)
        const level = Math.floor(Math.random() * 4 + 1)

        const device = new Device({
          uuid: uuidv4(),
          token: value.toString(8).slice(2),
          title: `Device ${i + 1}`,
          certificate: new Certificate(),
          tags: Array.from({ length }).map((_, i) => `tag ${i}`),
          access: new Access({ namespace: key, level, role }),
          enabled: true,
          config: new Struct({})
        })

        devices.set(device.uuid, device)
        return device
      }))
    }
  })

  service(DevicesService, {
    async list(request) {
      await new Promise((resolve) => setTimeout(resolve, 300))

      if (!request.namespace) {
        return new Devices({
          devices: Array.from(devicesByNs.values()).flat(1)
        })
      }

      return new Devices({
        devices: devicesByNs.get(request.namespace ?? '')
      })
    },
    get(request) {
      return devices.get(request.uuid) ?? request
    },
    create(request) {
      devices.set(request.namespace, new Device(request.device))
      const device = devices.get(request.device?.uuid ?? '')

      for (const devices of devicesByNs.values()) {
        devices.push(new Device(request.device))
      }

      return new CreateResponse({ device })
    },
    update(request) {
      devices.set(request.uuid, request)

      for (const devices of devicesByNs.values()) {
        const i = devices.findIndex(({ uuid }) => uuid === request.uuid)

        if (i !== -1) devices.splice(i, 1, request)
      }

      return devices.get(request.uuid) ?? request
    },
    patchConfig(request: Device): Device {
      const device = devices.get(request.uuid)
      if (!device) {
        throw new ConnectError("Device not found or not enough Access Rights", Code.NotFound)
      }

      device.config = request.config
      devices.set(request.uuid, device)
      return device
    
    },
    delete(request) {
      devices.delete(request.uuid)

      for (const devices of devicesByNs.values()) {
        const i = devices.findIndex(({ uuid }) => uuid === request.uuid)

        if (i !== -1) devices.splice(i, 1)
      }

      return new DeleteResponse()
    },
    join(request) {
      let node = nodes.get(request.node)?.find(
        ({ node }) => node.includes(request.join)
      )
      const { namespace } = devices.get(request.node)?.access ?? {}

      if (!nodes.has(request.node)) {
        nodes.set(request.node, [])
      }

      if (!node) {
        node = new Node({
          node: request.node,
          access: new Access({
            level: request.access,
            role: Role.SHARED,
            namespace
          }),
          edge: '',
          parent: ''
        })
        nodes.get(request.node)?.push(node)
      } else if (node.access) {
        node.access.level = request.access
      }

      return new Node(node)
    },
    joins(request) {
      return new Nodes({ nodes: nodes.get(request.uuid) })
    },
    async move(request) {
      const { namespaces } = await namespaceApi.list(new EmptyMessage())
      const namespace = namespaces.find(({ uuid }) => uuid === request.namespace)
      const device = devices.get(request.uuid) ?? new Device({ uuid: request.uuid })

      for (const [key, devices] of devicesByNs) {
        const i = devices.findIndex(({ uuid }) => uuid === request.uuid)

        if (i !== -1) {
          devicesByNs.get(key)?.splice(i, 1)
          break
        }
      }

      device.access = namespace?.access
      devicesByNs.get(request.namespace)?.push(device)
      return new EmptyMessage()
    },
    toggle(request) {
      return new Device(changeDevice('enabled', request.uuid, !request.enabled))
    },
    toggleBasic(request) {
      return new Device(changeDevice('basicEnabled', request.uuid, !request.basicEnabled))
    },
    makeDevicesToken(request) {
      const token = Math.random().toString(16).slice(2)

      for (const key of Object.keys(request.devices)) {
        changeDevice('token', key, token)
      }
      return new TokenResponse({ token })
    },
    getByToken(request) {
      return new Device(request)
    },
    getByFingerprint(request) {
      let device = new Device()
      for (const value of devices.values()) {
        const { fingerprint } = value.certificate ?? {}

        if (fingerprint === request.fingerprint) {
          device = value
        }
      }

      return new Device(device)
    }
  })
})
