        // Handle form submission
        document.getElementById("postForm").addEventListener("submit", function (e) {
            e.preventDefault(); // Prevent the default form submission

            // Get the values of title and body inputs
            var title = document.getElementById("title").value;
            var body = document.getElementById("body").value;

            // Get the feedback element
            var feedback = document.getElementById("feedback");

            // Check if both title and body are filled
            if (title.trim() === "" || body.trim() === "") {
                feedback.innerHTML = '<p style="color: red; font-weight: bold;">❌ Please fill in both Title and Body fields.</p>';
            } else {
                feedback.innerHTML = '<p style="color: green; font-weight: bold;">✅ Post created successfully!</p>';
            }

            // Clear the feedback message after a few seconds (adjust the time as needed)
            setTimeout(function () {
                feedback.innerHTML = '';
            }, 3000);
        });