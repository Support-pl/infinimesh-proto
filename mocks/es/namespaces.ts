import { Struct } from '@bufbuild/protobuf'
import { createPromiseClient, createRouterTransport } from '@connectrpc/connect'
import { AccountsService, NamespacesService } from '../../build/es/node/node_connect'
import { DeleteResponse, EmptyMessage } from '../../build/es/node/node_pb'
import { Access, Nodes, Node, Level, Role } from '../../build/es/node/access/access_pb'
import { Namespace, Namespaces, Plugin } from '../../build/es/node/namespaces/namespaces_pb'
import { Account, Accounts } from '../../build/es/node/accounts/accounts_pb.js'
import { transport as accTransport } from './accounts.ts'
import { v4 as uuidv4 } from 'uuid'

export const transport = createRouterTransport(({ service }) => {
  const accountsApi = createPromiseClient(AccountsService, accTransport)

  const namespaces = new Map<string, Namespace>()
  const accountsList = new Map<string, Account>()
  const accountsByNs = new Map<string, Account[]>()
  const nodes = new Map<string, Node[]>()

  const role = Math.floor(Math.random() * 2 + 1)
  const level = Math.floor(Math.random() * 4 + 1)

  for (const key of namespaces.keys()) {
    nodes.set(key, [0, 1].map((i) =>
      new Node({
        access: new Access({ namespace: key, level, role }),
        edge: `Namespases2Accounts/${key}-${i}`,
        parent: `Namespaces/${key}`,
        node: ''
      })
    ))
  }

  accountsApi.list(new EmptyMessage()).then(({ accounts }) => {
    accounts.forEach((account) => {
      const { namespace = '' } = account.access ?? {}

      if (accountsByNs.has(namespace)) {
        accountsByNs.get(namespace)?.push(account)
      } else {
        accountsByNs.set(namespace, [account])
      }

      accountsList.set(account.uuid, account)
    })
  })

  service(NamespacesService, {
    list() {
      if (namespaces.size > 0) {
        return new Namespaces({
          namespaces: Array.from(namespaces.values())
        })
      }

      return new Namespaces({
        namespaces: Array.from({ length: 10 }).map((_, i) => {
          const namespace = `${i}`
          const uuid = uuidv4()
          const role = Role.OWNER
          const level = Level.ADMIN

          const namespaceItem = new Namespace({
            uuid: namespace,
            title: `Namespace ${i + 1}`,
            plugin: new Plugin({ uuid, vars: {} }),
            access: new Access({ namespace, level, role }),
            config: new Struct({})
          })

          namespaces.set(namespace, namespaceItem)
          return namespaceItem
        })
      })
    },
    get(request) {
      return namespaces.get(request.uuid) ?? request
    },
    create(request) {
      namespaces.set(request.uuid, request)
      return namespaces.get(request.uuid) ?? request
    },
    update(request) {
      namespaces.set(request.uuid, request)
      return namespaces.get(request.uuid) ?? request
    },
    delete(request) {
      namespaces.delete(request.uuid)
      return new DeleteResponse()
    },
    join(request) {
      const account = accountsByNs.get(request.namespace)?.find(
        ({ uuid }) => uuid === request.account
      )

      if (account?.access) {
        account.access.level = request.access
      } else {
        const account = accountsList.get(request.account) ??
          new Account({ uuid: request.account })

        accountsByNs.get(request.namespace)?.push(account)
      }
      return new Accounts({ accounts: accountsByNs.get(request.namespace) })
    },
    joins(request) {
      return new Accounts({ accounts: accountsByNs.get(request.uuid) })
    },
    accessibles(request) {
      return new Nodes({ nodes: nodes.get(request.uuid) })
    },
    deletables(request) {
      return new Nodes({ nodes: nodes.get(request.uuid) })
    }
  })
})
