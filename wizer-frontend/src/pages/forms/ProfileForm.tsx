import React, {
  //useEffect, 
  useState
} from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Profile } from "@/models/Profile";
import Button from "@/components/ui/button";
import api from "@/services/api";




const ProfileForm = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const isNew = !id;

  const [profile, setProfile] = useState<Profile>({
    id: id ?? "",
    user_id: "",
    gender: "",
    birth_year: 0,
    height_cm: 0,
    weight_kg: 0,
    body_fat_percentage: "",
    goal: "",
  });

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()

    const newProfile: Profile = {
      user_id: profile.user_id,
      gender: profile.gender,
      birth_year: profile.birth_year,
      height_cm: profile.height_cm,
      weight_kg: profile.weight_kg,
      body_fat_percentage: profile.body_fat_percentage,
      goal: profile.goal,
    }

    try {

      let response = await api.profile.create(newProfile);
      if (response) {
        navigate("/profiles")
      }
    } catch (error) {
      console.log("Handle submit", profile)
    }
  }


  const handleChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setProfile((prev) => ({
      ...prev,
      [event.target.name]: event.target.value,
    }))
  }

  return (
    <>
      <div className="max-w-2xl mx-auto mt-10">
        <h1>{isNew ? "Create New Profile" : `Edit Profile ${id}`}</h1>
        <form onSubmit={handleSubmit} className="space-y-4">

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">User</label>
            <input
              type="text"
              name="user_id"
              value={profile.user_id}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Gender</label>
            <input
              type="text"
              name="gender"
              value={profile.gender}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Birth Year</label>
            <input
              type="number"
              name="birth_year"
              value={profile.birth_year}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Height(cm)</label>
            <input
              type="number"
              name="height_cm"
              value={profile.height_cm}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Weight (Kg)</label>
            <input
              type="number"
              name="weight_kg"
              value={profile.weight_kg}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Body Fat (%)</label>
            <input
              type="text"
              name="body_fat_percentage"
              value={profile.body_fat_percentage}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Goal</label>
            <textarea
              rows={5}
              cols={33}
              name="goal"
              value={profile.goal}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          {isNew ?
            <Button type="submit">Create Profile</Button> :
            <Button type="button" >Update Profile</Button>
          }
        </form>
      </div>


    </>

  )
}

export default ProfileForm;
