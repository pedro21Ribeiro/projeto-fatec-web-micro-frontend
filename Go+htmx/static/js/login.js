document.addEventListener("DOMContentLoaded", (event) => {
    document.body.addEventListener("htmx:beforeSwap", function (e) {
        if (e.detail.xhr.status === 422){

            e.detail.shouldSwap = true;
            e.detail.isError = false;
        }
    })
});


function test() {
    alert("test");
}