// import { render, screen, waitFor } from '@testing-library/vue'

// import { payoutsMock } from '../../../mocks/processMocks'
// import ProcessComponent from '../ProcessComponent.vue'

// jest.mock('request', () => ({
//   get: jest.fn()
// }))

// describe('Payout Component', () => {
//   beforeEach(() => {
//     jest
//       .spyOn(request, 'get')
//       .mockImplementation(jest.fn(() => Promise.resolve({ data: payoutsMock })))
//   })

//   test('it should load payout table after ', async () => {
//     let elements = []
//     const ui = render(ProcessComponent)

//     await waitFor(() => {
//       elements = ui.getAllByTestId('payout-collection-row')
//       expect(elements.length).toBe(11)
//     })
//   })
// })

import { vi, describe, it, expect, beforeEach } from 'vitest'

import { mount } from '@vue/test-utils'
import ProcessComponent from '../ProcessComponent.vue'
import request from '../../../internals/api'
import { processMock } from '../../../mocks/processMocks'

describe('HelloWorld', () => {
  beforeEach(() => {
    vi.spyOn(request, 'get').mockImplementation(vi.fn(() => Promise.resolve({ data: processMock })))
  })
  it('renders properly', () => {
    const wrapper = mount(ProcessComponent, {})
    expect(wrapper.text()).toContain('Hello Vitest')
  })
})
