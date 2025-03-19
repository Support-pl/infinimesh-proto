import { Timestamp } from '@bufbuild/protobuf'
import { createRouterTransport } from '@connectrpc/connect'
import { SessionsService } from '../../build/es/node/node_connect'
import { DeleteResponse } from '../..//build/es/node/node_pb';
import { Activity, Sessions } from '../../build/es/node/sessions/sessions_pb'
import { v4 as uuidv4 } from 'uuid';

function randomDate(start, end) {
    return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
}

export const transport = createRouterTransport(({ service }) => {
    service(SessionsService, {
        get() {
            return new Sessions({
                sessions: Array.from({ length: 10 }).map(() => {
                    const created = Timestamp.fromDate(randomDate(new Date(2022, 0, 1), new Date()))
                    const expires = Timestamp.fromDate(randomDate(created.toDate(), new Date()))
                    return { current: false, client: 'test' + uuidv4(), created, expires, id: uuidv4() }
                })
            })
        },
        getActivity() {
            const lastSeen = {}
            Array.from({ length: 10 }).map(() => {
                lastSeen[uuidv4()] = Timestamp.fromDate(randomDate(new Date(2022, 0, 1), new Date()))
            })

            return new Activity({ lastSeen })
        },

        revoke() {
            return new DeleteResponse()
        }
    })
})
