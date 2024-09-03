import List from "@/components/ui/list";
const renderProfileItem = (profile: { id: number; name: string; user: string }, column: string) => {
  switch (column) {
    case "Id":
      return (<p className="text-gray-700 font-semibold">{profile.id}</p>)
    case "Name":
      return (<p className="text-blue-500 font-semibold">{profile.name}</p>)
    case "User":
      return (<p className="text-blue-500 font-semibold">{profile.user}</p>)
    default:
      return "";
  };
}
// <div className="flex justify-between items-center">
//   <p className="text-gray-700">{profile.id}</p>
//   <p className="text-blue-500">{profile.name}</p>
// </div>

const ProfileList = () => {
  const columns = ["Id", "Name", "User"]
  const emptyItems = Array.from({ length: 20 }).map(() => ({}));
  const profiles = [
    { id: 1, name: "test 1", user: "user 1" },
    { id: 2, name: "test 2", user: "user 2" },
  ]

  const diplayedProfiles = [...profiles, ...emptyItems].slice(0, 20);



  return (
    <div className="flex flex-col">
      <h2 className="text-2xl font-bold mb-4">Profiles</h2>
      <div className="flex justify-center px-4 rounded-lg">
        <List
          items={diplayedProfiles}
          renderItem={renderProfileItem}
          columns={columns}
          className="border rounded-lg p-2"
        />
      </div>
    </div>
  )
}


export default ProfileList;



