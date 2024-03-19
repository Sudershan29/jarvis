import axios from 'axios';

export const getProfile = async (token) => {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/users/profile`, {
            headers: {
                Authorization: `Bearer ${token}`,
                "ngrok-skip-browser-warning": "69420"
            }
        });
        return response.data.user;
    } catch (error) {
        console.error('Error fetching profile:', error);
        return {};
    }
}

export const getCalendars = async (token) => {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/calendars`, {
            headers: {
                Authorization: `Bearer ${token}`,
                "ngrok-skip-browser-warning": "69420"
            }
        });
        return response.data.calendars;
    } catch (error) {
        console.error('Error fetching calendars:', error);
        return [];
    }
}

export const GoogleCalendarConnectFn = (token) => `${process.env.EXPO_PUBLIC_BACKEND_URL}/calendar/connect?token=${token}`