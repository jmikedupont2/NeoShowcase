import { User } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { client } from '/@/libs/api'
import { createMemo, createResource, createRoot } from 'solid-js'

const [users, { mutate: mutateUsers, refetch: refetchUsers }] = createResource(async () => {
  const getUsersRes = await client.getUsers({})
  return getUsersRes.users
})

export { users, mutateUsers, refetchUsers }

// keyにID, valueにユーザー情報を持つMap
// createRootを使わない場合``computations created outside a `createRoot` or `render` will never be disposed``の警告が出る
// see: https://www.solidjs.com/docs/latest#createroot
const usersMap = createRoot(() =>
  createMemo(() => {
    if (!users()) return new Map<string, User>()
    return new Map(users().map((user) => [user.id, user]))
  }),
)

export const userFromId = (id: string) => usersMap().get(id)
