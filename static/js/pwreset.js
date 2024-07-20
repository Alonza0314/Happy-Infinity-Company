document.addEventListener("DOMContentLoaded", function() {
    function checkPwresetNull(password, again) {
        return password !== "" && again !== "";
    }
    
    function checkPwresetPassword(password) {
        return password.length >= 4;
    }
    
    function checkPwresetAgain(password, again) {
        return password == again;
    }
    
    function checkPwreset() {
        var password = document.getElementById("password").value;
        var again = document.getElementById("password-again").value;
    
        if (!checkPwresetNull(password, again)) {
            alert("blank field\n\n欄位空白");
            return false;
        }
    
        if (!checkPwresetPassword(password)) {
            alert("password length must be greater than or equal to 4\nplease re-entern\n\n密碼長度需大於等於4，請重新輸入");
            return false;
        }
    
        if (!checkPwresetAgain(password, again)) {
            alert("password is different\nplease re-enter\n\n密碼不相同，請重新輸入");
            return false;
        }
        return true;
    }

    document.querySelector('.pwreset-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkPwreset()) {
            return;
        }
        this.submit();
    });
});