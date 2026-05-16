
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { VirenmonitoringSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await VirenmonitoringSDK.test()
    equal(null !== testsdk, true)
  })

})
