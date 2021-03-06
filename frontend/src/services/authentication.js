import { BehaviorSubject } from "rxjs";
import axios from "axios";
import jwt_decode from "jwt-decode";
import moment from "moment";

const currentUserSubject = new BehaviorSubject(JSON.parse(localStorage.getItem("currentUser")));

export const authenticationService = {
  login,
  logout,
  loginValid,
  currentUser: currentUserSubject.asObservable(),
  get currentUserValue() {
    return currentUserSubject.value;
  }
};

function login(username, password) {
  return axios
    .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/authentication`, {
      username: username,
      password: password
    })
    .then(res => {
      if (res.status === 200) {
        var decodedUser = jwt_decode(res.data);
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorage.setItem("currentUser", JSON.stringify(decodedUser));
        currentUserSubject.next(decodedUser);
        return decodedUser;
      }

      if (res.status === 401 || res.status === 403) {
        localStorage.removeItem("currentUser");
        currentUserSubject.next(null);
        return res.data;
      }
    })
    .catch(() => {
      return "User or password incorrrect!";
    });
}

function loginValid() {
  if (currentUserSubject && currentUserSubject.value) {
    return currentUserSubject.value.exp > moment().unix() ? true : false;
  } else return false;
}

function logout() {
  // remove user from local storage to log user out
  localStorage.removeItem("currentUser");
  currentUserSubject.next(null);
}
