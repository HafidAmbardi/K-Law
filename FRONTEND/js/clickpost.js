document.addEventListener("DOMContentLoaded", function() {
    const posts = document.querySelectorAll(".post");
    const filterContainer = document.querySelector(".filter-container");
    const sortByContainer = document.querySelector(".sort-by");
    const postContainer = document.querySelector(".post-container");
    const backButton = document.querySelector(".back-button");
    let prevScrollY = 0;
    let filteredPostsIndices = []; // Store indices of filtered posts before enlarging

    posts.forEach((post, index) => {
        // Store indices of filtered posts
        if (post.style.display === "none") {
            filteredPostsIndices.push(index);
        }

        post.addEventListener("click", function(event) {
            // Store current scroll position
            prevScrollY = window.scrollY;

            // Scroll to the top of the page
            window.scrollTo({
                top: 0,
                behavior: "smooth" // Optional: smooth scrolling effect
            });

            const moreSpan = post.querySelector("#more");
            const dotsSpan = post.querySelector("#dots");
            if (moreSpan) {
                moreSpan.style.display = "inline";
            }
            if (dotsSpan) {
                dotsSpan.style.display = "none";
            }

            // Show back button
            backButton.style.display = "block";

            // Hide filter buttons and Sort By container
            filterContainer.style.display = "none";
            sortByContainer.style.display = "none";

            // Store filtered posts indices
            filteredPostsIndices = [];
            posts.forEach((otherPost, otherIndex) => {
                if (otherPost !== post && otherPost.style.display === "none") {
                    filteredPostsIndices.push(otherIndex);
                }
            });

            // Hide other posts
            posts.forEach(otherPost => {
                if (otherPost !== post) {
                    otherPost.style.display = "none";
                }
            });

            // Enlarge clicked post
            post.classList.add("enlarged");
            post.classList.add("no-hover");

            // Add class to post container
            postContainer.classList.add("affected-container");

            // Prevent propagation of click events from the enlarged post section
            event.stopPropagation();

            // Add onclick attribute to the post itself
            post.setAttribute("onclick", "stopProp(event)");

            // Add onclick attribute to all elements inside the post
            const elementsInsidePost = post.querySelectorAll("*");
            elementsInsidePost.forEach(element => {
                element.setAttribute("onclick", "stopProp(event)");
            });

            // Add class to body
            document.body.classList.add("enlarged-body");
        });
    });

    // Handle back button click
    backButton.addEventListener("click", function() {
        // Hide back button
        backButton.style.display = "none";

        // Restore filter buttons and Sort By container
        filterContainer.style.display = "flex";
        sortByContainer.style.display = "flex";

        // Restore other posts
        posts.forEach((post, index) => {
            if (!filteredPostsIndices.includes(index)) {
                post.style.display = "block";
            }
        });

        // Shrink enlarged posts
        posts.forEach(post => {
            post.classList.remove("enlarged");
            post.classList.remove("no-hover");

            // Reset more and ... elements
            const moreSpan = post.querySelector("#more");
            const dotsSpan = post.querySelector("#dots");
            if (moreSpan) {
                moreSpan.style.display = "none";
            }
            if (dotsSpan) {
                dotsSpan.style.display = "inline";
            }

            // Remove onclick attribute from the post itself
            post.removeAttribute("onclick");

            // Remove onclick attribute from all elements inside the post
            const elementsInsidePost = post.querySelectorAll("*");
            elementsInsidePost.forEach(element => {
                element.removeAttribute("onclick");
            });
        });

        // Remove class from post container
        postContainer.classList.remove("affected-container");

        // Remove class from body
        document.body.classList.remove("enlarged-body");

        // Restore previous scroll position
        window.scrollTo({
            top: prevScrollY,
            behavior: "auto" // Optional: adjust as needed
        });
    });
});

function stopProp(event) {
    event.stopPropagation();
}