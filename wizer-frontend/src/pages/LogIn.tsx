import React, { useState } from "react";
import Button from "@/components/ui/button";

const LogIn = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  //const [passwordError, setPasswordError] = useState("");


  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

  }

  return (
    <div className="flex justify-center items-center h-screen">
      <form onSubmit={handleSubmit} className="w-full max-w-sm">
        <div className="mb-4">
          <label htmlFor="mb-4" className="block text-gray-700 font-bold mb-2">
            Email
          </label>
          <input
            type="email"
            id="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="mb-4" className="block text-gray-700 font-bold mb-2">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
          <Button type="submit">Log In</Button>
          <p className="text-center mt-2">
            Don't have an account? <a href="/signin" className="text-blue-500">Sign In</a>
          </p>
        </div>
      </form>

    </div>
  )

}

export default LogIn;
