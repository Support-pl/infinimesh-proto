import { Struct, Timestamp } from '@bufbuild/protobuf'
import { createPromiseClient, createRouterTransport } from '@connectrpc/connect'
import { DevicesService, ShadowService } from 'infinimesh-proto/build/es/node/node_connect'
import { QueryRequest } from 'infinimesh-proto/build/es/node/node_pb'
import { ConnectionState, GetResponse, Shadow, State } from 'infinimesh-proto/build/es/shadow/shadow_pb'
import { transport as devicesTransport } from './devices.ts'

export const transport = createRouterTransport(({ service }) => {
  const devicesApi = createPromiseClient(DevicesService, devicesTransport)
  const shadows = new Map<string, Shadow>()

  devicesApi.list(new QueryRequest()).then(({ devices }) => {
    devices.forEach(({ uuid }) => {
      const length = Math.floor(Math.random() * 6)
      const timestamp = new Timestamp({ seconds: BigInt(Math.round(Date.now() / 1000)) })

      shadows.set(uuid, new Shadow({
        connection: new ConnectionState({
          clientId: `${length}`, connected: true, timestamp
        }),
        desired: new State({ data: new Struct(), timestamp }),
        reported: new State({ data: new Struct(), timestamp }),
        device: uuid
      }))
    })
  })

  service(ShadowService, {
    get (request) {
      return new GetResponse({
        shadows: request.pool.map((uuid) => shadows.get(uuid)) as Shadow[]
      })
    },
    patch (request) {
      shadows.set(request.device, request)
      return new Shadow(request)
    },
    remove (request) {
      const result = shadows.get(request.device)

      shadows.delete(request.device)
      return new Shadow(result)
    },
    streamShadow (request) {
      return {
        [Symbol.asyncIterator]: () => ({
          next: async () => (request.devices.length > 0)
            ? { value: new Shadow(shadows.get(request.devices.pop() as string)), done: false }
            : { value: null, done: true }
        })
      }
    },
    streamShadowSync (request) {
      return {
        [Symbol.asyncIterator]: () => ({
          next: async () => (request.devices.length < 1)
            ? { value: new Shadow(shadows.get(request.devices.pop() as string)), done: false }
            : { value: null, done: true }
        })
      }
    }
  })
})
