import { Button } from '@/components/ui/button';
import { DatePicker, LocalizationProvider } from '@/components/ui/date-picker';
import { format } from 'date-fns';
import { useRouter } from 'next/router';
import { Controller, FormProvider, useForm } from 'react-hook-form';
import { useCreateSchedule } from './hooks/create-schedule';

type FormValue = {
  start_date: Date;
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
