
import { useAuth0 } from "@auth0/auth0-react";

  const userNavigation = [
    { name: 'Your profile', href: '/profile' },
    { name: 'Sign out' },
]

  interface MobileNavLoginBtnProps {
    isAuthenticated: boolean;
  }
  
  const MobileNavLoginBtn: React.FC<MobileNavLoginBtnProps> = ({ isAuthenticated }) => {
  const { user, loginWithRedirect, logout } = useAuth0();

    return (              
        <div>
        {isAuthenticated ? (
          <div className="py-6">
            <span className="ml-4 text-sm font-semibold leading-6 text-gray-900" aria-hidden="true">
            <img
                className="h-8 w-8 rounded-full bg-gray-50"
                src={user?.picture}
                alt=""
            />
                {user?.name}
            </span>
            {userNavigation.map((item) => (
                  <a
                    key={item.name}
                    onClick={() => item.name === 'Sign out' ? logout() : null}
                    href={item.href}
                    className="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                  >
                    {item.name}
                  </a>
                ))}
            </div>
        ) : (
        <div className="py-6">
            <a
                onClick={() => loginWithRedirect()}
                className="-mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
            >
                Log in
            </a>
        </div>
        )}
      </div>
    )
}

export default MobileNavLoginBtn;