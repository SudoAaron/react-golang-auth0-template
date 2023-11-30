"use client";
import React, { useEffect, useState } from "react";
import { Dialog } from "@headlessui/react";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";
import NavLoginBtn from "../Buttons/LoginBtn/NavLoginBtn";
import ProfileDropdown from "../profileDropdown/ProfileDropdown";
import MobileNavLoginBtn from "../Buttons/LoginBtn/MobileLoginBtn";
import { useAuth0 } from "@auth0/auth0-react";
import axios from "axios";

interface NavigationItem {
  name: string;
  href: string;
}

const companyName = "React-Golang-Auth0";

const navigation: NavigationItem[] = [
  { name: "Home", href: "/" },
  { name: "Public", href: "/public" },
  { name: "Private", href: "/private" },
  { name: "Admin", href: "/admin" },
];

export default function NavBar() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [isLoggingIn, setIsLoggingIn] = useState(false);

  const { isAuthenticated, getAccessTokenSilently, isLoading } = useAuth0();

  const handleLogin = () => {
    setIsLoggingIn(true);
  };

  useEffect(() => {
    const initUser = async () => {
      const accessToken = await getAccessTokenSilently({
        authorizationParams: {
          audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
          scope: "openid profile email read:messages",
        },
      });
      axios({
        url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/init_user",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }).catch((err: Error) => {
        console.error(err);
      });
    };
    if (isAuthenticated) {
      initUser();
    }
  }, [isLoggingIn]);

  return (
    <header className="bg-white">
      <nav
        className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8"
        aria-label="Global"
      >
        <div className="flex items-center gap-x-12">
          <a href="#" className="-m-1.5 p-1.5">
            <span className="sr-only">{companyName}</span>
            <img
              className="h-8 w-auto"
              src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
              alt=""
            />
          </a>
          <div className="hidden lg:flex lg:gap-x-12">
            {navigation.map((item) => (
              <a
                key={item.name}
                href={item.href}
                className="text-sm font-semibold leading-6 text-gray-900"
              >
                {item.name}
              </a>
            ))}
          </div>
        </div>
        <div className="flex lg:hidden">
          <button
            type="button"
            className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
            onClick={() => setMobileMenuOpen(true)}
          >
            <span className="sr-only">Open main menu</span>
            <Bars3Icon className="h-6 w-6" aria-hidden="true" />
          </button>
        </div>
        <div className="hidden lg:flex">
          {!isLoading && isAuthenticated ? (
            <ProfileDropdown />
          ) : (
            !isLoading && <NavLoginBtn handleLogin={handleLogin} />
          )}
        </div>
      </nav>
      <Dialog
        as="div"
        className="lg:hidden"
        open={mobileMenuOpen}
        onClose={() => setMobileMenuOpen(false)}
      >
        <div className="fixed inset-0 z-10" />
        <Dialog.Panel className="fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
          <div className="flex items-center justify-between">
            <a href="#" className="-m-1.5 p-1.5">
              <span className="sr-only">{companyName}</span>
              <img
                className="h-8 w-auto"
                src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                alt=""
              />
            </a>
            <button
              type="button"
              className="-m-2.5 rounded-md p-2.5 text-gray-700"
              onClick={() => setMobileMenuOpen(false)}
            >
              <span className="sr-only">Close menu</span>
              <XMarkIcon className="h-6 w-6" aria-hidden="true" />
            </button>
          </div>
          <div className="mt-6 flow-root">
            <div className="-my-6 divide-y divide-gray-500/10">
              <div className="space-y-2 py-6">
                {navigation.map((item) => (
                  <a
                    key={item.name}
                    href={item.href}
                    className="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                  >
                    {item.name}
                  </a>
                ))}
              </div>
              <div>
                <MobileNavLoginBtn handleLogin={handleLogin} />
              </div>
            </div>
          </div>
        </Dialog.Panel>
      </Dialog>
    </header>
  );
}
