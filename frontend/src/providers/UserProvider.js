import React, { createContext, useState, useEffect } from "react";
const UserContext = createContext([{}, () => {}]);

const UserProvider = props => {
  const [state, setState] = useState({ selectedUser: null });

  // Get all users
  useEffect(() => {
  }, []); // Empty array mean, exacute once this effect

  return <UserContext.Provider value={[state, setState]}>{props.children}</UserContext.Provider>;
};

export { UserContext, UserProvider };
