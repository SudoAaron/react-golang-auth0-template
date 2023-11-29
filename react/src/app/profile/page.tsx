"use client"
import LoadingSpinner from "@/components/loading/LoadingSpinner";
import { withAuthenticationRequired } from "@auth0/auth0-react";
import { useAuth0 } from "@auth0/auth0-react";

export const Profile = () => {

const {
    user,
} = useAuth0();

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