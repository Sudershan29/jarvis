import axios from 'axios';

export const getEvents = async (token, startDate, endDate) => {
    try {
        const params = {};
        if (startDate) {
            params.startDate = startDate;
        }
        if (endDate) {
            params.endDate = endDate;
        }

        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/calendars/events`, {
            params,
            headers: {
                Authorization: `Bearer ${token}`,
                "ngrok-skip-browser-warning": "69420"
            }
        });
        return response.data.events;
    } catch (error) {
        console.error('Error fetching events:', error);
        return [];
    }
}
