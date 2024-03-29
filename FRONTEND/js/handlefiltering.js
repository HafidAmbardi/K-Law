document.addEventListener('DOMContentLoaded', function() {
    // Function to handle filter button clicks
    function handleFilterButtonClick() {
        const selectedCompanies = Array.from(document.querySelectorAll('.companyButton.selected')).map(button => button.textContent.trim());
        const selectedViolations = Array.from(document.querySelectorAll('.violationButton.selected')).map(button => button.textContent.trim());

        // Show posts matching selected tag(s) and apply fade-in effect
        document.querySelectorAll('.post').forEach(post => {
            const companyTags = Array.from(post.querySelectorAll('.company-tag p')).map(tag => tag.textContent.trim());
            const violationTags = Array.from(post.querySelectorAll('.violation-tag p')).map(tag => tag.textContent.trim());

            // Check if post contains all selected companies and any selected violations
            const companiesMatch = selectedCompanies.length === 0 || selectedCompanies.some(company => companyTags.includes(company));
            const violationsMatch = selectedViolations.length === 0 || selectedViolations.some(violation => violationTags.includes(violation));

            if (companiesMatch && violationsMatch) {
                post.style.opacity = '1'; // Set opacity to 1 to show post
                post.style.display = 'block'; // Ensure post is displayed
            } else {
                post.style.opacity = '0'; // Set opacity to 0 to hide post
                post.style.display = 'none'; // Hide post
            }
        });
    }

    // Add event listeners to filter buttons
    document.querySelectorAll('.violationButton, .companyButton').forEach(button => {
        button.addEventListener('click', () => {
            button.classList.toggle('selected'); // Toggle selected class
            handleFilterButtonClick(); // Handle filter button click
        });
    });
});
