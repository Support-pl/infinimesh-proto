import { createRouterTransport } from '@connectrpc/connect'
import { AccountsService } from "../../build/es/node/node_connect"
import { Access } from "../../build/es/node/access/access_pb"
import { Account, Accounts, CreateResponse, CreateRequest } from "../../build/es/node/accounts/accounts_pb"
import { DeleteResponse, EmptyMessage, SetCredentialsResponse, TokenResponse } from '../../build/es/node/node_pb'
import { v4 as uuidv4 } from 'uuid'

function generateAccount(i = 0) {
  const namespace = `${Math.floor(Math.random() * 10)}`
  const role = Math.floor(Math.random() * 2 + 1)
  const level = Math.floor(Math.random() * 4 + 1)

  const accountItem = new Account({
    uuid: uuidv4(),
    title: `Account ${i + 1}`,
    access: new Access({ namespace, level, role }),
    enabled: true,
    config: {}
  })

  return accountItem
}

export const transport = createRouterTransport(({ service }) => {
  service(AccountsService, {
    list() {
      return new Accounts({
        accounts: Array.from({ length: 10 }).map((_, i) => generateAccount(i))
      })
    },
    get(account: Account) {
      return { ...generateAccount(), uuid: account.uuid }
    },
    create(request: CreateRequest) {
      return new CreateResponse({ account: { ...request, title: "test", config: {}, defaultNamespace: "1", enabled: true, access: {}, uuid: uuidv4() } })
    },
    update(account: Account) {
      return new Account({ ...account })
    },
    toggle(account: Account) {
      return new Account({ ...account, enabled: !account.enabled })
    },
    delete() {
      return new DeleteResponse()
    },
    move() {
      return new EmptyMessage()
    },
    setCredentials() {
      return new SetCredentialsResponse()
    },
    token() {
      return new TokenResponse({ token: uuidv4() })
    }
  })
})
