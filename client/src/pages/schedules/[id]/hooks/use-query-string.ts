import { useRouter } from "next/router"

export const useQueryString = (key: string) =>{
	const router = useRouter()
	const value =	router.query[key]
 	if (!value) return
	if (typeof value !== 'string') return
	return	value
}