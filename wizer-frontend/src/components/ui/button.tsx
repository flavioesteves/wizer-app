import { ReactNode, HTMLAttributes } from "react";
import { cn } from "@/lib/utils";

type ButtonProps = HTMLAttributes<HTMLButtonElement> & {
  variant?: "primary" | "secondary" | "green" | "red"; // Define allowed variants
  size?: "sm" | "md" | "lg"; // Define allowed sizes
  type?: "button" | "submit" | "reset";
  children: ReactNode; // Represent any child element
}

const Button: React.FC<ButtonProps> = ({ variant = "primary", size = "md", children, className, ...otherProps }) => {
  // Define button styles based on variant and size
  const buttonClasses = cn(`
    text-white py-2 px-4 rounded-md
    ${variant === "primary" ? "bg-blue-500 hover:bg-blue-700" : ""}
    ${variant === "secondary" ? "bg-gray-400 hover:bg-gray-500" : ""}
    ${variant === "green" ? "bg-green-300 hover:bg-green-400" : ""} 
    ${variant === "red" ? "bg-red-300 hover:bg-red-400" : ""} 
    ${size === "sm" ? "px-2 py-1" : ""}
    ${size === "lg" ? "px-6 py-3" : ""}
  `, className
  );

  return (
    <button className={buttonClasses} {...otherProps} > {children}</button >
  );
};

export default Button;
