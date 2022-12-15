import {shallowMount} from "@vue/test-utils";
import App from "@/App";
import API from "@/api";
import flushPromises from "flush-promises";

jest.mock('@/api')

describe('App.vue', () => {
    // App component exist check
    test("App component should exist", () => {
        // Mounted the App component
        const wrapper = shallowMount(App)
        // If component exists it will return true
        expect(wrapper.exists()).toBeTruthy()
    })

    // Checks the text box input exist
    test(' Text box input should exist', () => {
        // Mounted the App component
        const wrapper = shallowMount(App)
        // Find the given selector element
        const textbox = wrapper.find('.text-box')

        // If input exists it will return true
        expect(textbox.exists()).toBeTruthy()
    })

    // Checks the add button exist and button text should equals to "Add"
    test('Button should exist and render button text correctly', () => {
        // Mounted the App component
        const wrapper = shallowMount(App)
        // Find the given selector element
        let button = wrapper.find('.add-button')

        // If button exists it will return true
        expect(button.exists()).toBeTruthy()
        // Expect that button text equals to "Add"
        expect(button.text()).toEqual('Add')
    })

    // When user clicks the button, button should call the "addTask" function
    test('Button functionality check', async () => {
        // Spy on App component's addTask method
        const spyMethod = jest.spyOn(App.methods, 'addTask')

        // Mount the component and find the given element and click it
        await shallowMount(App).find('.add-button').trigger('click')

        // Expect that "addTask" method call
        expect(spyMethod).toHaveBeenCalled()

    })

    // Checks the add function success
    test('Add a task should successfully', async () => {
        // Local variable created
        const localThis = {
            tasks: [
                {
                    id: 1,
                    title: "test"
                },
                {
                    id: 2,
                    title: "test2"
                }
            ],
            newTask: "burcu",
            taskId: null,
        }

        // Mock the API getTasks method , it should return tasks array
        API.getTasks.mockResolvedValue(localThis.tasks)
        // App component method call with local data
        await App.methods.addTask.call(localThis)

        // Wait all awaits
        await flushPromises();

        // Expect that local tasks length to be 3 , added "burcu" task
        expect(localThis.tasks.length).toBe(3)
        // Expect that last tasks title to be "burcu"
        expect(localThis.tasks[localThis.tasks.length - 1].title).toBe("burcu")
    })
})