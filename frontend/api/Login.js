
import axios from 'axios';

export async function userLogin(userEmail, password) {
    try {
        const response = await axios.post(`${process.env.EXPO_PUBLIC_BACKEND_URL}/login`, {
            email: userEmail,
            password: password
        }, {
            "ngrok-skip-browser-warning": "69420"
        });

        return { success: true, token: response.data.token }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export async function userGoogleLogin() {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/google/signin`, {
            "ngrok-skip-browser-warning": "69420"
        });
        return { success: true, token: response.data.token }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export const GoogleLoginURL = () => `${process.env.EXPO_PUBLIC_BACKEND_URL}/google/signin`

export const GoogleCallbackURL = () => `${process.env.EXPO_PUBLIC_BACKEND_URL}/google/callback`