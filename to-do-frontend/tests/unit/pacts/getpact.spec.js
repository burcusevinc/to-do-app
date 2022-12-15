import {pactWith} from 'jest-pact';
import { Matchers } from '@pact-foundation/pact';
const { eachLike, like } = Matchers
import {API} from "@/api";

pactWith({
    consumer: "UserInterface", // Consumer name
    provider: "TodoInterface", // Provider name
}, provider => {
    describe("Todo Application", () => {
        let api
        beforeEach(() => {
            // Mocked backend/provider
            api = new API(provider.mockService.baseUrl)
        })
        test('Get task list', async () => {
            // Interaction -> define request + response
            await provider.addInteraction({ //mock service
                // Provider state
                state: 'get task list successfully',
                // Pact file test name
                uponReceiving: 'a request not empty for task list',
                // Request object, method and path
                withRequest: {
                    method: 'GET',
                    path: '/api/v1/tasks',
                },
                // Response object, status code and header
                willRespondWith: {
                    status: 200,
                    headers: {
                        'Content-Type': 'application/json; charset=UTF-8',
                    },
                    // Http body
                    body: eachLike({
                        id: like(1),
                        title: like("drink water"),
                    })
                }
            })
            // Mock provider's method it returns body
            const res = await api.getTasks()
            // Checks if the id of the 0th data is equals to 1
            expect(res[0].id).toEqual(1)
        })
    })
})