const {Given, When, Then} = require("cucumber");
const openUrl = require("../support/action/openUrl");
const assert = require("assert");

Given(/^Empty ToDo list$/, async function () {
    //it opens the given url
    await openUrl.call(this, '/')
    //find the given selector
    const selected = await this.page.$(".text-box")
    //it returns the selector's value on json format
    const textbox = await (await (selected).getProperty("value")).jsonValue()

    //assert that selector value's length equals to zero
    assert.equal(textbox.length, 0)
});
When(/^User write "([^"]*)" to text box$/, async function (inputText) {
    //find the given element selector
    const textbox = await this.page.$(".text-box");
    //focus the textbox element
    await textbox.focus();
    //type the "buy a milk" to textbox element
    await textbox.type(inputText);
    //wait
    await this.page.waitForTimeout(3000)
});
When(/^and User click to add button$/, async function () {
    //find the given element selector
    const button = await this.page.$(".add-button");
    //click the button
    await button.click();
    //wait
    await this.page.waitForTimeout(3000)
});
Then(/^User should see "([^"]*)" text in ToDo list$/, async function (inputText) {
    const selector = "#app"
    // Find the elements in root container(#app), which one text includes the "buy a milk"
    const task = await this.page.$$eval (
        //first parameter
        selector,
        //second parameter
        async (items, inputText) => {
            const task = items.find(item => item.textContent.includes(inputText))
            //if task is undefined, return false
            return !!task
        },
        //third parameter
        inputText
    )

    // Assert that elements returns true
    assert.strictEqual(task, true)
    // Wait
    await this.page.waitForTimeout(2000)
});