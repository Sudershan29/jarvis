
import React, { createContext, useState, useEffect } from 'react'
import AsyncStorage from "@react-native-async-storage/async-storage"
import { userLogin } from "../api/Login"

export const AuthContext = createContext()

export const AuthProvider = ({children}) => {
    const [isLoading, setIsLoading] = useState(false)
    const [userToken, setUserToken] = useState(null)

    useEffect(() => {
        checkLoginStatus()
    }, []);

    async function login (email, password) {
        setIsLoading(true)
        const token = await userLogin(email, password)
        setUserToken(token)
        AsyncStorage.setItem('@LoginStore:userToken', token)
        setIsLoading(false)
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

    function isLoggedIn() {
        return userToken !== null
    }

    return (
        <AuthContext.Provider value={{ userToken, isLoading, isLoggedIn, login, logout}}>
            {children}
        </AuthContext.Provider>
    )
}