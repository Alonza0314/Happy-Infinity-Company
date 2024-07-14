function Language() {
    var ci = document.getElementById("contact-info");
    var cien = document.getElementById("contact-info-en");
    
    if (ci.style.display === "none") {
        ci.style.display = "block";
        cien.style.display = "none";
    } else {
        ci.style.display = "none";
        cien.style.display = "block";
    }
}