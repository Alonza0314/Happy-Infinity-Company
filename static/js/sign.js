let signin = document.getElementById('signin');
let signup = document.getElementById('signup');
let form_box = document.getElementsByClassName('form-box')[0];
let signin_box = document.getElementsByClassName('signin-box')[0];
let signup_box = document.getElementsByClassName('signup-box')[0];

function isDesktop() {
    return window.innerWidth > 820;
}

if (isDesktop()) {
    signin.addEventListener('click', () => {
        form_box.style.transform = 'translateX(0%)';
        signup_box.classList.add('hidden');
        signin_box.classList.remove('hidden');
    });
    
    signup.addEventListener('click', () => {
        form_box.style.transform = 'translateX(87%)';
        signin_box.classList.add('hidden');
        signup_box.classList.remove('hidden');
    });
}
else {
    signin.addEventListener('click', () => {
        signup_box.classList.add('hidden');
        signin_box.classList.remove('hidden');
    });
    
    signup.addEventListener('click', () => {
        signin_box.classList.add('hidden');
        signup_box.classList.remove('hidden');
    });
}

document.addEventListener("DOMContentLoaded", function() {
    function checkSignupNull(username, email, password, passwordAgain) {
        return username !== "" && email !== "" && password !== "" && passwordAgain !== ""
    }
    
    function checkSignupUsername(username) {
        var regex = /^(?=.*\d)[a-z\d]+$/;
        return regex.test(username);
    }
    
    function checkSignupEmail(email) {
        var regex = /\S+@\S+\.\S+/;
        return regex.test(email);
    }
    
    function checkSignupPassword(password) {
        return password.length >= 4;
    }
    
    function checkSignupPasswordAgain(password, passwordAgain) {
        return password == passwordAgain;
    }
    
    function checkSignup() {
        var username = document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var password = document.getElementById("password").value;
        var passwordAgain = document.getElementById("password-again").value;
    
        if (!checkSignupNull(username, email, password, passwordAgain)) {
            alert("blank field\n\n欄位空白");
            return false;
        }
    
        if (!checkSignupUsername(username)) {
            alert("username can only and must contain lowercase English and numbers\nplease re-enter\n\n使用者名稱僅能且必須包含小寫英文及數字，請重新輸入");
            return false;
        }
    
        if (!checkSignupEmail(email)) {
            alert("email format error\nplease re-enter\n\n電子郵件格式錯誤，請重新輸入");
            return false;
        }
    
        if (!checkSignupPassword(password)) {
            alert("password length must be greater than or equal to 4\nplease re-entern\n\n密碼長度需大於等於4，請重新輸入");
            return false;
        }
    
        if (!checkSignupPasswordAgain(password, passwordAgain)) {
            alert("password is different\nplease re-enter\n\n密碼不相同，請重新輸入");
            return false;
        }
    
        return true;
    }
    
    function checksigninNull() {
        var username = document.getElementById("signin-username").value;
        var password = document.getElementById("signin-password").value;
    
        if (username == "" || password == "") {
            alert("blank field\n\n欄位空白");
            return false;
        }
        return true;
    }
    
    document.querySelector('.signup-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkSignup()) {
            return;
        }
        this.submit();
    });

    document.querySelector('.signin-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checksigninNull()) {
            return;
        }
        this.submit();
    });
});

function getUrlParameter(name) {
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    var results = regex.exec(location.search);
    return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
}

var signupCondition= getUrlParameter('signup');
if (signupCondition === 'true') {
    alert('signup success\nplease re-signin\n\n註冊成功\n請重新登入');
} else {
    if (signupCondition !== '') {
        alert('fail: ' + signupCondition);
    }
}

var signinCondition = getUrlParameter('signin');
if (signinCondition !== '') {
    alert('fail: ' + signinCondition);
}

var resetpwCondition = getUrlParameter('resetpw');
if (resetpwCondition !== '') {
    alert(resetpwCondition);
}