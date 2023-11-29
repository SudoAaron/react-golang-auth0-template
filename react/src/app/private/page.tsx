"use client"
import React, {useState, useEffect} from 'react';
import { useAuth0, withAuthenticationRequired } from "@auth0/auth0-react";
import LoadingSpinner from "@/components/loading/LoadingSpinner";
import axios from 'axios';

export const Private = () => {
  const [message, setMessage] = useState("");
  const {
    isAuthenticated,
    getAccessTokenSilently,
  } = useAuth0();

  useEffect(() => {
    const accessPage = async () => {
      const domain = process.env.NEXT_PUBLIC_AUTH0_DOMAIN
      const accessToken = await getAccessTokenSilently({
        authorizationParams: {
          audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
          scope: "openid profile email read:messages",
        },
      });

      axios({
        url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/private",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      })
      .then((res) => {
        setMessage(res.data.message)
      })
      .catch((err) => {
        setMessage(err.message)
      });
    }
    accessPage();
    }, [getAccessTokenSilently, isAuthenticated])

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <p>{message}</p>
      </div>
    </main>
  )
}

export default withAuthenticationRequired(Private, {
  onRedirecting: () => <LoadingSpinner />,
});
