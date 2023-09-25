import { Button } from '@/components/ui/button';
import { DatePicker, LocalizationProvider } from '@/components/ui/date-picker';
import { format } from 'date-fns';
import { useRouter } from 'next/router';
import { Controller, FormProvider, useForm } from 'react-hook-form';
import { useCreateSchedule } from './hooks/create-schedule';

type FormValue = {
  start_date: Date;
};

type ResponseBody = {
  schedule: {
    id: number;
    user_id: string;
    start_date: string;
  };
};

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

const CreateMain = () => {
  const router = useRouter();
  const { createSchedule, isResponseBody, isErrorBody } = useCreateSchedule();
  const methods = useForm<FormValue>();
  const { handleSubmit } = methods;

  const onClick = async ({ start_date }: FormValue) => {
    const reqBody = { start_date: format(start_date, 'yyyy-MM-dd') };
    const res = await createSchedule(reqBody);

    if (isResponseBody(res)) {
      router.push(`/schedules/${res.schedule.id}`);
    }

    if (isErrorBody(res)) {
      alert(res.message);
    }
  };

  return (
    <FormProvider {...methods}>
      <Controller
        name={'start_date'}
        render={({ field }) => (
          <LocalizationProvider>
            <DatePicker {...field} />
          </LocalizationProvider>
        )}
      />
      <Button onClick={handleSubmit(onClick)}>作成</Button>
    </FormProvider>
  );
};

export default CreateMain;
