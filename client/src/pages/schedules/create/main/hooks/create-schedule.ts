import { serverDomain } from "@/api/domain";
import { HttpMethods } from "@/api/methods";

type RequestBody = {
	start_date: string
}

type ResponseBody = {
  schedule: {
    id: number;
    user_id: string;
    start_date: string;
  };
};

type ErrorBody = {
	message: string
}

export const useCreateSchedule = () => {
	const createSchedule = async (req: RequestBody) => {
    const res = await fetch(`${serverDomain}/schedules`, {
      method: HttpMethods.POST,
      body: JSON.stringify(req),
    });
    const resJSON = await res.json();
		console.log(resJSON)

		if (isResponseBody(resJSON) || isErrorBody(resJSON)) {
      return resJSON
	}		else {
      throw new Error('createSchedule is failed: 不適切なフォーマット');
    }
	}

	const isResponseBody = (body: unknown): body is ResponseBody => {
		if (!body || typeof body !== 'object' || !('schedule' in body)) {
			return false;
		}

		const schedule = body.schedule;
		if (!schedule || typeof schedule !== 'object') {
			return false;
		}
		if (!('id' in schedule) || typeof schedule.id !== 'number') {
			return false;
		}
		if (!('user_id' in schedule) || typeof schedule.user_id !== 'string') {
			return false;
		}
		if (!('start_date' in schedule) || typeof schedule.start_date !== 'string') {
			return false;
		}
		return true;
	};

	const isErrorBody = (body: unknown): body is ErrorBody => {
		if(!body || typeof body !== 'object' ) return false
		if('message' in body && typeof body.message === 'string') {
			return true
		}
		return false
	}

	return {createSchedule, isResponseBody, isErrorBody}
}