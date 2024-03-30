document.addEventListener("DOMContentLoaded", function() {
    const posts = document.querySelectorAll(".post");
    const filterContainer = document.querySelector(".filter-container");
    const sortByContainer = document.querySelector(".sort-by");
    const postContainer = document.querySelector(".post-container");
    const backButton = document.querySelector(".back-button");
    let prevScrollY = 0;
    let filteredPostsIndices = [];

    posts.forEach((post, index) => {
        if (post.style.display === "none") {
            filteredPostsIndices.push(index);
        }

        post.addEventListener("click", function(event) {
            prevScrollY = window.scrollY;

            window.scrollTo({
                top: 0,
                behavior: "smooth" 
            });

            const moreSpan = post.querySelector("#more");
            const dotsSpan = post.querySelector("#dots");
            if (moreSpan) {
                moreSpan.style.display = "inline";
            }
            if (dotsSpan) {
                dotsSpan.style.display = "none";
            }

            backButton.style.display = "block";
            filterContainer.style.display = "none";
            sortByContainer.style.display = "none";

            filteredPostsIndices = [];
            posts.forEach((otherPost, otherIndex) => {
                if (otherPost !== post && otherPost.style.display === "none") {
                    filteredPostsIndices.push(otherIndex);
                }
            });

            posts.forEach(otherPost => {
                if (otherPost !== post) {
                    otherPost.style.display = "none";
                }
            });

            post.classList.add("enlarged");
            post.classList.add("no-hover");

            postContainer.classList.add("affected-container");

            event.stopPropagation();

            post.setAttribute("onclick", "stopProp(event)");

            const elementsInsidePost = post.querySelectorAll("*");
            elementsInsidePost.forEach(element => {
                element.setAttribute("onclick", "stopProp(event)");
            });

            document.body.classList.add("enlarged-body");
        });
    });

    backButton.addEventListener("click", function() {
        backButton.style.display = "none";

        filterContainer.style.display = "flex";
        sortByContainer.style.display = "flex";

        posts.forEach((post, index) => {
            if (!filteredPostsIndices.includes(index)) {
                post.style.display = "block";
            }
        });

        posts.forEach(post => {
            post.classList.remove("enlarged");
            post.classList.remove("no-hover");

            const moreSpan = post.querySelector("#more");
            const dotsSpan = post.querySelector("#dots");
            if (moreSpan) {
                moreSpan.style.display = "none";
            }
            if (dotsSpan) {
                dotsSpan.style.display = "inline";
            }

            post.removeAttribute("onclick");

            const elementsInsidePost = post.querySelectorAll("*");
            elementsInsidePost.forEach(element => {
                element.removeAttribute("onclick");
            });
        });

        postContainer.classList.remove("affected-container");

        document.body.classList.remove("enlarged-body");

        window.scrollTo({
            top: prevScrollY,
            behavior: "auto" 
        });
    });
});

function stopProp(event) {
    event.stopPropagation();
}