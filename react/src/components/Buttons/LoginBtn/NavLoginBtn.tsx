import { useAuth0 } from "@auth0/auth0-react";

export default function NavLoginBtn() { 
  const { loginWithRedirect } = useAuth0();

    return (              
      <div>
          <a 
            className="text-sm font-semibold leading-6 text-gray-900"
            onClick={() => loginWithRedirect()}
          >
            Log in <span aria-hidden="true">&rarr;</span>
          </a>
        </div>
    )
}