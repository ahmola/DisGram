package dev.ahmola.photogram.user_service.error;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class UserNotFoundException extends ApiException {
    public UserNotFoundException() {
        super(ErrorCode.ENTITY_NOT_FOUND);
        log.error("User Not Found");
    }
}
