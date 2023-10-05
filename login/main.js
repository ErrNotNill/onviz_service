const container = document.querySelector(".container"),
  pwShowHide = document.querySelectorAll(".showHidePw"),
  pwFields = document.querySelectorAll(".password"),
  signUp = document.querySelector(".signup-link"),
  login = document.querySelector(".login-link");

//   js code to show/hide password and change icon
pwShowHide.forEach(eyeIcon =>{
  eyeIcon.addEventListener("click", ()=>{
    pwFields.forEach(pwField =>{
      if(pwField.type ==="password"){
        pwField.type = "text";

        pwShowHide.forEach(icon =>{
          icon.classList.replace("uil-eye-slash", "uil-eye");
        })
      }else{
        pwField.type = "password";

        pwShowHide.forEach(icon =>{
          icon.classList.replace("uil-eye", "uil-eye-slash");
        })
      }
    })
  })
})

// js code to appear signup and login form
signUp.addEventListener("click", ( )=>{
  container.classList.add("active");
});
login.addEventListener("click", ( )=>{
  container.classList.remove("active");
});

let password = document.getElementById('password').value;
let passwordConfirm = document.getElementById('passwordConfirm').value
console.log(password)
console.log(passwordConfirm)


function myFunction() {
  let email = "name@email.com";
  let password = `pass`;
  console.log(document.getElementById('password').value);
  if (document.getElementById('password').value === password){
    console.log('success');
  } else {
    console.log(document.getElementById('password').value);
    console.log('Failure');
    console.log(password);
    console.log(document.getElementById('password').value);
  }
}


console.log(password)
console.log(passwordConfirm)
if (passwordConfirm !== password) {

  console.error("Password confirmation failed");
} else {
  fetch("http://localhost:9090/auth_page", {
    method: "POST",
    body: JSON.stringify({
      name: document.getElementsByName('username'),
      email: document.getElementsByName('email'),
      password: document.getElementsByName('password')
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then(r => errorOccurred());
}

//name,email,password,passwordConfirm

