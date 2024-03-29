document.addEventListener('DOMContentLoaded', function() {
    // Function to handle filter button clicks
    function handleFilterButtonClick() {
        const selectedCompanies = Array.from(document.querySelectorAll('.companyButton.selected')).map(button => button.textContent.trim());
        const selectedViolations = Array.from(document.querySelectorAll('.violationButton.selected')).map(button => button.textContent.trim());

        document.querySelectorAll('.post').forEach(post => {
            const companyTags = Array.from(post.querySelectorAll('.company-tag p')).map(tag => tag.textContent.trim());
            const violationTags = Array.from(post.querySelectorAll('.violation-tag p')).map(tag => tag.textContent.trim());

            const companiesMatch = selectedCompanies.length === 0 || selectedCompanies.some(company => companyTags.includes(company));
            const violationsMatch = selectedViolations.length === 0 || selectedViolations.some(violation => violationTags.includes(violation));

            if (companiesMatch && violationsMatch) {
                post.style.opacity = '1'; 
                post.style.display = 'block'; 

                post.style.opacity = '0'; 
                post.style.display = 'none'; 
            }
        });
    }

    document.querySelectorAll('.violationButton, .companyButton').forEach(button => {
        button.addEventListener('click', () => {
            button.classList.toggle('selected'); 
            handleFilterButtonClick(); 
        });
    });
});
