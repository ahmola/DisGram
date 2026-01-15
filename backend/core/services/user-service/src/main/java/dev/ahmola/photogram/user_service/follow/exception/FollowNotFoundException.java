package dev.ahmola.photogram.user_service.follow.exception;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class FollowNotFoundException extends ApiException {
    public FollowNotFoundException(){
        super(ErrorCode.ENTITY_NOT_FOUND);
        log.error("Follow Not Found");
    }
}
