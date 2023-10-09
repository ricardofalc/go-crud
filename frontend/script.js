document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("postForm");

    form.addEventListener("submit", function (event) {
        event.preventDefault();

        const formData = new FormData(form);
        const jsonData = {};

        formData.forEach(function (value, key) {
            jsonData[key] = value;
        });

        const requestOptions = {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(jsonData),
        };

        fetch("/posts", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                console.log("Post created:", data);
                // You can add logic to handle the response here
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    });
});
