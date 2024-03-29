document.querySelectorAll(".companyButton"). forEach(companyButton => {
    companyButton.addEventListener("click", () => companyButton.classList.toggle("blue-selected"));
});

document.querySelectorAll(".violationButton"). forEach(violationButton => {
    violationButton.addEventListener("click", () => violationButton.classList.toggle("red-selected"));
});