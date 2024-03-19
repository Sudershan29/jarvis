
import React, { createContext, useState, useEffect } from 'react'
import AsyncStorage from "@react-native-async-storage/async-storage"
import { userLogin, userGoogleLogin } from "../api/Login"
import { showMessage, hideMessage } from "react-native-flash-message";

export const AuthContext = createContext()

export const AuthProvider = ({children}) => {
    const [isLoading, setIsLoading] = useState(false)
    const [userToken, setUserToken] = useState(null)
    const [connectedToCalendar, setConnectedToCalendar] = useState(false)

    useEffect(() => {
        checkLoginStatus()
    }, []);

    async function login (email, password) {
        // setIsLoading(true)
        const response = await userLogin(email, password)
        if (!response.success){
            showMessage({
                message: response.message,
                type: "error",
            })
            // setIsLoading(false);
            return;
        }
        setUserToken(response.token)
        AsyncStorage.setItem('@LoginStore:userToken', response.token)
        setIsLoading(false)
    }

    async function loginGoogleSuccess (token) {
        setIsLoading(true)
        setUserToken(token)
        AsyncStorage.setItem('@LoginStore:userToken', token)
        setIsLoading(false)
    }

    async function loginGoogle () {
        // setIsLoading(true)
        const response = await userGoogleLogin()
        if (!response.success){
            setIsLoading(false);
            showMessage({
                message: response.message,
                type: "error",
            });
            return;
        }

        setUserToken(response.token)
        AsyncStorage.setItem('@LoginStore:userToken', token)
        // setIsLoading(false)
    }

    function logout () {
        setIsLoading(true)
        setUserToken(null)
        AsyncStorage.removeItem('@LoginStore:userToken')
        setIsLoading(false)
    }
    
    async function checkLoginStatus() {
        try{
            setIsLoading(true)
            let userToken = await AsyncStorage.getItem('@LoginStore:userToken');
            setUserToken(userToken)
            setIsLoading(false)
        }catch(e){
            console.log("cannot load userTokens")
        }
    }

    function calendarConnectSuccess(whatTodoAfter){
        setConnectedToCalendar(true)
        whatTodoAfter()
    }

    function isLoggedIn() {
        return userToken !== null
    }

    return (
        <AuthContext.Provider value={{ userToken, isLoading, connectedToCalendar, isLoggedIn, login, logout, loginGoogle, loginGoogleSuccess, calendarConnectSuccess }}>
            {children}
        </AuthContext.Provider>
    )
}