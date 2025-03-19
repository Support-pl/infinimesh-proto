import { createRouterTransport } from '@connectrpc/connect'
import { PluginsService } from '../../build/es/plugins/plugins_connect'
import { Plugin, Plugins } from '../../build/es/plugins/plugins_pb';
import { v4 as uuidv4 } from 'uuid';

export const transport = createRouterTransport(({ service }) => {
    service(PluginsService, {
        list() {
            return new Plugins({
                pool: Array.from({ length: 10 }).map(() => {
                    return { public: false, uuid: uuidv4(), description: 'test' + Math.random().toString().slice(2), title: 'test' + uuidv4(), }
                })
            })
        },
        create(plugin) {
            return new Plugin({ ...plugin, uuid: uuidv4() })
        },
        update(plugin) {
            return plugin
        }
    })
})
