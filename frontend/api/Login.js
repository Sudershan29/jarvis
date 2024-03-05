
import axios from 'axios';
import { BACKEND_URL } from "react-native-dotenv"

export async function userLogin(userEmail, password) {
    try {
        const response = await axios.post(`${BACKEND_URL}/login`, {
            email: userEmail,
            password: password
        });

        return { success: true, token: response.data.token }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export async function userGoogleLogin() {
    try {
        const response = await axios.get(`${BACKEND_URL}/google/signin`, { headers: {'Access-Control-Allow-Origin': '*'}});
        return { success: true, token: response.data.token }
    } catch (error) {
        return { success: false, message: error.message }
    }
}