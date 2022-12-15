import {pactWith} from 'jest-pact';
import {Matchers} from '@pact-foundation/pact';

const {eachLike, like} = Matchers
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
        test('Post a task', async () => {
            // Interaction -> define request + response
            await provider.addInteraction({ //mock service
                // Provider state
                state: 'post a task successfully',
                // Pact file test name
                uponReceiving: 'add not empty task',
                // Request object, method and path
                withRequest: {
                    method: 'POST',
                    path: '/api/v1/tasks',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: {
                        id: like(1),
                        title: like("drink water")
                    }
                },
                // Response object, status code and header
                willRespondWith: {
                    status: 201,
                }
            })
            const task = {
                id: 1,
                title: "drink water"
            }
            await api.createTasks(task)
        })
    })
})