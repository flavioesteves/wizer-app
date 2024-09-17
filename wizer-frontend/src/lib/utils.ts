import { twMerge } from "tailwind-merge";
import { clsx, ClassValue } from "clsx";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}


export function createArrayFromObject(obj: any): string[] {
  return Object.keys(obj).map((key) => {
    return `${key}`
  });
}
