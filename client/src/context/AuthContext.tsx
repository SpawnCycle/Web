import { createContext } from "react";

export const AuthContext = createContext<{
  isLoggedIn: boolean;
  setIsLoggedIn: (value: boolean) => void;
}>({
  isLoggedIn: false,
  setIsLoggedIn: () => {},
});
