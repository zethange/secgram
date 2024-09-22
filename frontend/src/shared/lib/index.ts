export * from './type'
export * from './gradient'

export const func = () => {
  return new Promise(async (res, rej) => {
    res(undefined)
    rej()
  })
}