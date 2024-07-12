document.addEventListener('DOMContentLoaded', function() {
    var languageSwitcherZH = document.getElementById('languageSwitcherZH');
    var languageSwitcherEN = document.getElementById('languageSwitcherEN');
    
    languageSwitcherZH.addEventListener('click', function(event) {
        event.preventDefault();
        topbarLanguage();
        bottombarLanguage();
        Language();
    });
    
    languageSwitcherEN.addEventListener('click', function(event) {
        event.preventDefault();
        topbarLanguage();
        bottombarLanguage();
        Language();
    });
});