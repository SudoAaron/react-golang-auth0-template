"use client"
import React from "react";
import { useEffect, useState } from "react";
import LoadingSpinner from "@/components/loading/LoadingSpinner";
import { withAuthenticationRequired } from "@auth0/auth0-react";
import { useAuth0 } from "@auth0/auth0-react";
import { useRouter } from 'next/navigation'
import axios from 'axios';

export const Profile = () => {
const router = useRouter()
const [roles, setRoles] = useState([]);
const { user, getAccessTokenSilently, isAuthenticated } = useAuth0();

const getRoles = async () => {
  const accessToken = await getAccessTokenSilently({
    authorizationParams: {
      audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
      scope: "openid profile email read:messages",
    },
  });
  axios({
    url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/get_roles",
    method: "GET",
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  })
  .then((res) => {
    setRoles(res.data.roles)
  })
  .catch((err) => {
    if (err.response && err.response.status !== 200) {
        router.push("/")
    }
  });
}

useEffect(() => {
  getRoles();
  }, [getAccessTokenSilently, isAuthenticated])

  const handleBecomeAdmin = async function(){
    const accessToken = await getAccessTokenSilently({
      authorizationParams: {
        audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
        scope: "openid profile email read:messages",
      },
    });
    axios({
      url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/set_role",
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    })
    .then((res) => {
      getRoles();
    })
    .catch((err) => {
      console.error(err)
    });
  }

  return (
    <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">

    <main className="px-4 py-16 sm:px-6 lg:flex-auto lg:px-0 lg:py-20">
    <div className="mx-auto max-w-2xl space-y-16 sm:space-y-20 lg:mx-0 lg:max-w-none">
      <div>
        <h2 className="text-base font-semibold leading-7 text-gray-900">Profile</h2>

        <dl className="mt-6 space-y-6 divide-y divide-gray-100 border-t border-gray-200 text-sm leading-6">
          <div className="pt-6 sm:flex">
            <dt className="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">Full name</dt>
            <dd className="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto">
              <div className="text-gray-900">{user?.name}</div>
            </dd>
          </div>
          <div className="pt-6 sm:flex">
            <dt className="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">Email address</dt>
            <dd className="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto">
              <div className="text-gray-900">{user?.email}</div>
            </dd>
          </div>
          <div className="pt-6 sm:flex">
            <dt className="font-medium text-gray-900 sm:w-64 sm:flex-none sm:pr-6">Roles</dt>
            <dd className="mt-1 flex justify-between gap-x-6 sm:mt-0 sm:flex-auto">
              <div className="text-gray-900">
              {roles.map((role, index) => (
                <React.Fragment key={role}>
                  {role}
                  {index !== roles.length - 1 ? ', ' : ''}
                </React.Fragment>
              ))}
                </div>
            </dd>
          </div>
          <button
        type="button"
        className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        onClick={() => {handleBecomeAdmin()}}
      >
        Become Admin
      </button>
        </dl>
      </div>
      </div>
      </main>
        </div>
    </div>
  )
}

export default withAuthenticationRequired(Profile, {
  onRedirecting: () => <LoadingSpinner />,
});