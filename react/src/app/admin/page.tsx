"use client"
import React, {useState, useEffect} from 'react';
import { useAuth0, withAuthenticationRequired } from "@auth0/auth0-react";
import LoadingSpinner from "@/components/loading/LoadingSpinner";
import { useRouter } from 'next/navigation'
import axios from 'axios';

export const Admin = () => {
  const router = useRouter()
  const [message, setMessage] = useState("");
  const {
    isAuthenticated,
    getAccessTokenSilently,
  } = useAuth0();

  useEffect(() => {
    const accessPage = async () => {
      const accessToken = await getAccessTokenSilently({
        authorizationParams: {
          audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
          scope: "openid profile email read:messages",
        },
      });
      axios({
        url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/admin",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      })
      .then((res) => {
        setMessage(res.data.message)
      })
      .catch((err) => {
        if (err.response && err.response.status !== 200) {
            router.push("/")
        }
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

export default withAuthenticationRequired(Admin, {
  onRedirecting: () => <LoadingSpinner />,
});
