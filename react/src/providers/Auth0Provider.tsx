"use client"
import React, { ReactNode } from "react";
import { Auth0Provider } from '@auth0/auth0-react';
import { useRouter } from 'next/navigation'


type AppProviderProps = {
    children: ReactNode;
  };
  
export const AppProvider: React.FC<AppProviderProps> = ({ children }) => {
    const auth0Domain = process.env.NEXT_PUBLIC_AUTH0_DOMAIN || "";
    const auth0ClientId = process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID || "";

    const router = useRouter()

    const onRedirectCallback = (appState: any) => {
      router.push(
        appState && appState.returnTo ? appState.returnTo : window.location.pathname
      );
    };

    return (
    <Auth0Provider
        domain={auth0Domain}
        clientId={auth0ClientId}
        onRedirectCallback={onRedirectCallback}
        authorizationParams={{
          audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
          redirect_uri: "http://localhost:3000/",
          scope: "openid profile email read:messages",
        }}
    >
    {children}
  </Auth0Provider>
    )
}
