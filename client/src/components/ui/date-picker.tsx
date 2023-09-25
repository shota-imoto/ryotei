import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { DatePicker as MuiDatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider as MuiLocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import enUs from 'date-fns/locale/en-US';
import { ComponentProps, ReactNode, forwardRef } from 'react';

type Props = ComponentProps<typeof MuiDatePicker>;

export const DatePicker = forwardRef<any, Props>(function DatePicker(props, ref) {
  return <MuiDatePicker ref={ref} {...props} />;
});

type ProviderProps = {
  children: ReactNode;
};
export const LocalizationProvider = ({ children }: ProviderProps) => (
  <MuiLocalizationProvider adapterLocale={enUs} dateAdapter={AdapterDateFns}>
    {children}
  </MuiLocalizationProvider>
);
