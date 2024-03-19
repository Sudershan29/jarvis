
import React, { useContext, useState } from "react";
import {
    StyleSheet,
    View,
    TextInput,
    Button,
} from "react-native";
// import GoogleButton from 'react-google-button'
import { AuthContext } from "../../context/AuthContext"
import * as WebBrowser from 'expo-web-browser';
import { makeRedirectUri, useAuthRequest, ResponseType } from 'expo-auth-session';
import { GoogleLoginURL, GoogleCallbackURL }  from "../../api/Login"
/*

TODO: Add Input validation
TODO: Figure out how to give error popups

*/
// WebBrowser.maybeCompleteAuthSession({
//     showInRecents: true,
// });

// Endpoint
// const discovery = {
//     authorizationEndpoint: GoogleLoginURL(),
// };

export default function LoginScreen({navigation}) {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const { login } = useContext(AuthContext)

    // const [request, response, promptAsync] = useAuthRequest(
    //     {
    //         responseType: ResponseType.Token,
    //         clientId: 'CLIENT_ID',
    //         scopes: [],
    //         redirectUri: makeRedirectUri({
    //             // Ensure the scheme matches your expo application scheme
    //             scheme: 'your.app',
    //             useProxy: true,
    //         }),
    //     },
    //     discovery
    // );

    // React.useEffect(() => {
    //     if (response && response.type === 'success') {
    //         console.log(response)
    //         const token = response.params.access_token;
    //         // Handle the successful authentication here
    //     }
    // }, [response]);

    function NavigateToLogin(){
        window.location.href = GoogleLoginURL();
    }

    return (
        <View style={styles.container}>
            <View style={styles.inputView}>
                <TextInput
                    style={styles.TextInput}
                    placeholder="Email."
                    placeholderTextColor="#003f5c"
                    onChangeText={(email) => setEmail(email)}
                />
            </View>
            <View style={styles.inputView}>
                <TextInput
                    style={styles.TextInput}
                    placeholder="Password."
                    placeholderTextColor="#003f5c"
                    secureTextEntry={true}
                    onChangeText={(password) => setPassword(password)}
                />
            </View>

            <Button
                title="Login"
                onPress={() => login(email, password)}
            />

            <Button onPress={NavigateToLogin} title="Login with Google" />  

        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: "center",
        justifyContent: "center",
    },
    inputView: {
        borderRadius: 30,
        width: "70%",
        height: 45,
        marginBottom: 20,
        alignItems: "center",
    },
    TextInput: {
        height: 50,
        flex: 1,
        padding: 10,
        marginLeft: 20,
    }
});
