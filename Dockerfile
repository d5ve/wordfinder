FROM perl:5.34-slim-bullseye AS base

WORKDIR /app

FROM base AS build

RUN apt update \
    && apt install -y make gcc wbritish \
    && cpanm Dancer2 JSON::XS \
    && apt remove -y make gcc \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /usr/lib/gcc

FROM base AS app

# Try to copy the minimal files and dirs from the build image.
# Perl is spread around the place though, so it's tricky.
COPY --from=build /etc/dictionaries-common /etc/dictionaries-common
COPY --from=build /usr/lib /usr/lib
COPY --from=build /usr/local/bin /usr/local/bin
COPY --from=build /usr/local/lib/perl5/site_perl /usr/local/lib/perl5/site_perl
COPY --from=build /usr/share /usr/share

# Proper SIG handling.
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

# In-container debugging.
# RUN apt update && install -y vim procps

COPY wordfinder.psgi .

CMD ["/usr/local/bin/plackup", "--listen", ":80", "--app", "/app/wordfinder.psgi"]
