import { useAuth0 } from "@auth0/auth0-react";

interface NavLoginBtnProps {
  handleLogin: () => void;
}

const NavLoginBtn: React.FC<NavLoginBtnProps> = ({handleLogin}) => { 
  const { loginWithPopup, isAuthenticated } = useAuth0();

  const handleLoginClick = async () => {
    try {
      await loginWithPopup({authorizationParams: {}})
      handleLogin();
     } catch(e) {
      console.error(e)
     }
  };

    return (              
      <div>
          <a 
            className="text-sm font-semibold leading-6 text-gray-900"
            onClick={() => handleLoginClick()}
          >
            Log in <span aria-hidden="true">&rarr;</span>
          </a>
        </div>
    )
}

export default NavLoginBtn;