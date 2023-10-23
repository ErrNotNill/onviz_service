<template>
  <div class="container" :class="{ active: isSignup }">
    <div class="forms">
      <div class="form login">
        <span class="title">Onviz</span>


          <div class="input-field">
            <input v-model="loginEmail" type="text" placeholder="Enter your email" required>
            <i class="uil uil-envelope icon"></i>
          </div>
          <div class="input-field">
            <input v-model="loginPassword" type="password" class="password" placeholder="Enter your password" required>
            <i class="uil uil-lock icon"></i>
            <i class="uil uil-eye-slash showHidePw" @click="togglePasswordVisibility('loginPassword')"></i>
          </div>

          <div class="checkbox-text">
            <div class="checkbox-content">
              <input type="checkbox" id="logCheck">
              <label for="logCheck" class="text">Remember me</label>
            </div>

            <a href="#" class="text">Forgot password?</a>
          </div>

          <div class="input-field button">
            <input type="button" value="Login" @click="signinUser">
          </div>


        <div class="login-signup">
          <span class="text">Not a member?
            <a href="#" class="text signup-link" @click="switchToSignup">Signup Now</a>
          </span>
        </div>
      </div>

      <!-- Registration Form -->
      <div class="form signup">
        <span class="title">Registration</span>


          <div class="input-field" name="username" id="username">
            <input v-model="signupName" type="text" placeholder="Enter your name" required>
            <i class="uil uil-user"></i>
          </div>
          <div class="input-field" name="email" id="email">
            <input v-model="signupEmail" type="text" placeholder="Enter your email" required>
            <i class="uil uil-envelope icon"></i>
          </div>
          <div class="input-field" name="password" id="password">
            <input v-model="signupPassword" type="password" class="password" placeholder="Create a password" required>
            <i class="uil uil-lock icon"></i>
          </div>
          <div class="input-field" name="passwordConfirm" id="passwordConfirm">
            <input v-model="signupPasswordConfirm" type="password" class="password" placeholder="Confirm a password" required>
            <i class="uil uil-lock icon"></i>
            <i class="uil uil-eye-slash showHidePw" @click="togglePasswordVisibility('signupPasswordConfirm')"></i>
          </div>

          <div class="checkbox-text">
            <div class="checkbox-content">
              <input type="checkbox" id="termCon">
              <label for="termCon" class="text">I accepted all terms and conditions</label>
            </div>
          </div>

          <div class="input-field button">
            <input type="button" value="Signup" @click="signupUser">
          </div>


        <div class="login-signup">
          <span class="text">Already a member?
            <a href="#" class="text login-link" @click="switchToLogin">Login Now</a>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const isSignup = ref(false);
const loginEmail = ref('');
const loginPassword = ref('');
const signupName = ref('');
const signupEmail = ref('');
const signupPassword = ref('');
const signupPasswordConfirm = ref('');

function togglePasswordVisibility(passwordField) {
  const field = eval(passwordField);
  field.type = field.type === 'password' ? 'text' : 'password';
}

function validateEmail(email) {
  // Basic email format validation
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
}

function validatePassword(password) {
  // Password should contain at least 6 characters, a symbol, and both upper and lower case letters
  const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&(){}])[A-Za-z\d!@#$%^&(){}]{6,}$/;
  return passwordRegex.test(password);
}

function signinUser() {
  // Create an object with the user's login data
  const loginData = {
    email: loginEmail.value,
    password: loginPassword.value,
  };

  // Define the URL of your server where you want to send the login data
  const loginUrl = 'http://localhost:9090/login_page'; // Replace with your server URL

  // Send a POST request to the server
  fetch(loginUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(loginData),
  })
    .then(response => {
      if (response.status === 200) {
        // Login successful, get the token
        return response.text(); // Assuming the response is a SHA-256 token string
      } else {
        // Login failed, log response status and response data
        console.error('Login failed. Status:', response.status);
        return response.text().then(text => {
          console.error('Response data:', text);
          throw new Error('Login failed');
        });
      }
    })
    .then(token => {
      // Handle the response token
      console.log('Received token:', token);
      // Store the token in session storage
      sessionStorage.setItem('token', token);

      // Redirect to the desired URL
      window.location.href = 'http://localhost:5173/'; // Replace with the URL you want to redirect to
    })

}


function signupUser() {
  // Client-side validation
  if (signupPassword.value !== signupPasswordConfirm.value) {
    alert('Password confirmation failed.');
    return;
  }

  if (
    signupName.value.length < 6 ||
    !validateEmail(signupEmail.value) ||
    !validatePassword(signupPassword.value)
  ) {
    alert('Invalid input. Please check your details.');
    return;
  }

  // Create an object with the user's registration data
  const userData = {
    username: signupName.value,
    email: signupEmail.value,
    password: signupPassword.value,
  };

  // Define the URL of your server where you want to send the registration data
  const registrationUrl = 'http://localhost:9090/auth_page'; // Replace with your server URL

  // Send a POST request to the server
  fetch(registrationUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(userData),
  })
    .then(response => {
      if (response.status === 200) {
        // Registration successful
        return response.json();
      } else {
        // Log the error response for debugging
        console.error('Registration failed. Status:', response.status);
        return response.text().then(text => {
          console.error('Response data:', text);
          throw new Error('Registration failed');
        });
      }
    })
    .then(data => {
      // Handle the response data (if applicable)
      console.log('Server response:', data);
      if (data.message === 'Registration successful') {
        // Handle a successful registration
        alert('Registration successful!');
      } else {
        // Handle other responses from the server if needed
        if (data.message === 'User exist') {
          alert('User exists, please change the email address');
        } else {
          alert('Registration failed...');
        }
      }
    })
    .catch(error => {
      // Handle errors, e.g., registration failure
      console.error('Error:', error);
      alert('Registration failed...');
    });
}


function switchToSignup() {
  isSignup.value = true;
}

function switchToLogin() {
  isSignup.value = false;
}
</script>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');

*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Poppins', sans-serif;
}

body{
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #4070f4;
}

.container{
  position: relative;
  max-width: 430px;
  width: 100%;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin: 0 20px;
}

.container .forms{
  display: flex;
  align-items: center;
  height: 440px;
  width: 200%;
  transition: height 0.2s ease;
}


.container .form{
  width: 50%;
  padding: 30px;
  background-color: #fff;
  transition: margin-left 0.18s ease;
}

.container.active .login{
  margin-left: -50%;
  opacity: 0;
  transition: margin-left 0.18s ease, opacity 0.15s ease;
}

.container .signup{
  opacity: 0;
  transition: opacity 0.09s ease;
}
.container.active .signup{
  opacity: 1;
  transition: opacity 0.2s ease;
}

.container.active .forms{
  height: 600px;
}
.container .form .title{
  position: relative;
  font-size: 27px;
  font-weight: 600;
}

.form .title::before{
  content: '';
  position: absolute;
  left: 0;
  bottom: 0;
  height: 3px;
  width: 30px;
  background-color: #4070f4;
  border-radius: 25px;
}

.form .input-field{
  position: relative;
  height: 50px;
  width: 100%;
  margin-top: 30px;
}

.input-field input{
  position: absolute;
  height: 100%;
  width: 100%;
  padding: 0 35px;
  border: none;
  outline: none;
  font-size: 16px;
  border-bottom: 2px solid #ccc;
  border-top: 2px solid transparent;
  transition: all 0.2s ease;
}

.input-field input:is(:focus, :valid){
  border-bottom-color: #4070f4;
}

.input-field i{
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  color: #999;
  font-size: 23px;
  transition: all 0.2s ease;
}

.input-field input:is(:focus, :valid) ~ i{
  color: #4070f4;
}

.input-field i.icon{
  left: 0;
}
.input-field i.showHidePw{
  right: 0;
  cursor: pointer;
  padding: 10px;
}

.form .checkbox-text{
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 20px;
}

.checkbox-text .checkbox-content{
  display: flex;
  align-items: center;
}

.checkbox-content input{
  margin-right: 10px;
  accent-color: #4070f4;
}

.form .text{
  color: #333;
  font-size: 14px;
}

.form a.text{
  color: #4070f4;
  text-decoration: none;
}
.form a:hover{
  text-decoration: underline;
}

.form .button{
  margin-top: 35px;
}

.form .button input{
  border: none;
  color: #fff;
  font-size: 17px;
  font-weight: 500;
  letter-spacing: 1px;
  border-radius: 6px;
  background-color: #4070f4;
  cursor: pointer;
  transition: all 0.3s ease;
}

.button input:hover{
  background-color: #265df2;
}

.form .login-signup{
  margin-top: 30px;
  text-align: center;
}
</style>

