import auth from "./auth";
import user from "./user";


export default {
  getUsers: user.getAll,
  registerUser: user.register,
  loginUser: auth.signIn,
  validateSession: auth.validateSession,
}
